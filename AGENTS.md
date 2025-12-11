# PenguinTunes

This is a desktop music player app for linux based OSs built using Wails framework
The logic is handled by Go and UI is crafted using VueJS

## Frontend Conventions

- scripts and folder names should be in kebab-case style
- component names should PascalCase
- use setup script (TS) + composition API
- pinia stores also also should be written as setup store
- stores are defined in ./src/lib/stores
- inside components, the script tag HAS TO be on the top and style tag at the end
- write styles as tailwind as much as possible
- always use `@/` path alias in import statements
- if a specific style is not supported by tailwind, write styles inside `<style scoped>` tag
