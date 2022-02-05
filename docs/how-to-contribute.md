# How To contribute

The site will be available at <TBD>. This repository will always be public.

# Making a proposal

1. Fork this repository
2. Add a file to the `docs/proposals` directory with your feature proposal and raise a PR on this repository.
- use the format `docs/proposals/format.md`

# Maintainer review of idea

1. If the maintainers agree to this idea (which they probably will) they will merge this Pr.
2. This document will be the basis of discussion between the maintainers and you on how to develop this feature.

# Development
1. Once the document is approved it will be merged into main and you can continue developing
   on your fork.
2. You along with folks who would like to help can start developing the feature on your fork.

# Maintainer review
1. Once you are ready raise your PR. Here is more info on how to work on a fork.
   https://gist.github.com/Chaser324/ce0505fbed06b947d962
2. If you add new `db` migrations, please make sure you update the `schema.sql` file before you raise the PR.
3. In your final PR make sure that the proposal is moved to `docs/rfcs` from `docs/proposals`
   (`docs/rfcs` has only the merged features)

# Merge
1. Once we merge your PR into master CI/CD should kick in to deploy to production.
   
# Things Left To Do

- [ ] Deploy front end somewhere
- [ ] Setup the CI/CD
  - [DONE] Api CI/CD

# Notes

The idea behind this process is to lay a framework to make communication easier
and, more importantly, written down.
