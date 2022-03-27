# Twitter Login Flow

_view on large screen or with wrap off_

_Mobile flow_

TODO: Fix this for oauth2

**mobile**                          **tribist.com**                             **twitter.com**
    |     /users/login?source=twitter     |                                           |
    --------------->----------------------+                                           |
    |                                     |                                           |
    |                                     |                                           |
    | 302 to `/i/oauth2/authorize?client_id,redirect_uri`                             |               <---- redirect-uri has `scope`, `state`, `challenge` 
    +--------------<-----------------------                                           |
    |                                     |                                           |
    |                                     |                                           |
    |                `GET  /i/oauth2/authorize?client_id, redirect_uri`               |
    -------------------------------------->-------------------------------------------+
    |                                     |                                           |
    |                `302 /users/authorize/twitter?oauth_token,oauth-verifier`        |
    +-------------------------------------<--------------------------------------------
    |                                     |                                           |
    |   `GET /users/authorize/twitter?code,state`                                     |
    --------------->----------------------+                                           |
    |                                     |                                           |
    |                                     |    `POST /2/oauth2/token`                 |
    |                                     -------------------->-----------------------+                <---- passes `client_id`,`code`,`challenge`,`redirect_uri`
    |                                     |                                           |
    |                                     |                                           |
    |                                     +-------------------<------------------------                <---- returns `tokens`
    |                                     |                                           |
    |    `302 found`                      |                                           |
    +--------------<-----------------------                                           |
    |                                     |                                           |
    |                                     |                                           |
    |                                     |                                           |
    |                                     |                                           |
    |                                     |                                           |
    |                                     |                                           |
    |                                     |                                           |
    |                                     |                                           |
    |                                     |                                           |
    |                                     |                                           |
    |                                     |                                           |
    |                                     |                                           |
