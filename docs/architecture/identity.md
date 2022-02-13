# Identity

## description

Users can login using 

1. Google
2. Facebook
3. Twitter

We will store the following in our database
1. user's email id
2. the authentication source [ Google, Facebook, Twitter ]

## Overview 

The user has one unique `uuid` in the system. 
The user table looks like this

```
CREATE TYPE user_role AS ENUM ('user', 'admin');

CREATE TABLE users (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    handle TEXT NOT NULL UNIQUE
    profile TEXT -- url to photo
    role user_role
);
```

Each user may have multiple login methods

```
CREATE TYPE login_source AS ENUM ('google', 'twitter', 'facebook');

CREATE TABLE login_methods (
    id uuid default uuid_generate_v4(),
    user_id uuid NOT NULL,
    source login_source NOT NULL,
    attributes jsonb default '{}' -- source specific attributes
);
```

## User Tokens

We will use jwt tokens for login verification.

jwt token will contain

```
{
"user_id": <>,
"role": ["user"|"admin"],
}
```

On successful login, we will return two tokens 
1. access token
2. refresh token

*logout flow*

On logout we place the `access_token` and `refresh_token` in the `invalid_tokens` table.
On login or refresh we check the passed token exists in `invalid_tokens`, if so we forbid
the action.

```
CREATE TABLE invalid_tokens (
    token TEXT PRIMARY KEY
)
```
## APIs

*GET /users/login?source=*
    
    headers: None
    
    query: 
        source: [twitter|facebook|google]
        
    triggers the login flow for the source mentioned by redirecting somewhere.

*GET /users/authorize/twitter*

    query:
        `oauth_token`
        `oauth-verifier`
        
    returns:
        {
            "access_token": <access-token>,
            "refresh_token": <refresh-token>
        }
        
    At the end of the login flow on the login source we tell the server that 
    we are done and successful.

*GET /users/profile*

    headers:
        X-PROFILE-ACCESS-TOKEN: <access token> 

    returns:
        {
            "id":
            "name":
            "profile":
            "user_role":
        }
        
*POST /users/refresh*

    headers:
        X-PROFILE-REFRESH-TOKEN: <refresh token>
    
    returns:
        {
            "access_token": <new access token>,
            "refresh_token": <new refresh token> // may not change
        }
        
*POST /users/logout*

    headers:
        X-PROFILE-ACCESS-TOKEN: <access token>
    
    Invalidates the access token. Relogin requires the `refresh_token`
    or relogin flow.
    
## Use cases

- [ ] As a user I want to be able to create an account using Twitter
- [ ] As a user I want to be able to create an account using Google
- [ ] As a user I want to be able to create an account using Facebook
- [ ] As a user I want to be able to logout
- [ ] As a user I want to link my accounts together
