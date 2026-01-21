## The Prettier + Eslint issues

The problem is that we need to run Prettier as well as Eslint,
but plugin for breaking tailwindcss classes to multiple lines is only available
in Eslint. That means that Prettier will try to join the lines and Eslint will
split them afterwards, meaning both will report changes even if in the end
nothing was actually changed. Therefore, we need to run them with `|| true`
to ensure that they won't block the git hook and let on to the pre-commit to fail
only if there are changes in the files after both Prettier & Eslint has run.
