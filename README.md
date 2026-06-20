# DevLocal

DevLocal is a developer tool that allows teams to define local-only repository customizations that:

1. Make a project runnable on a developer machine.
2. Can be shared with the team through source control.
3. Never accidentally get committed.
4. Work across languages and frameworks.
5. Require no changes to existing Git workflows.

The goal is to eliminate README instructions such as:

* Change port X to Y
* Disable mypy locally
* Update docker-compose for your machine
* Add local credentials
* Modify logging settings

and replace them with:

```bash
devlocal apply
```