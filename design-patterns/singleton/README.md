# Singleton Design Pattern
A unique instance of a type in an entire program.

Provides a single instance of an object/garuntees no duplicates.
- First call, instance is created
- Aftewards, reused

## Objective
Used when...
1. Single, shared value, some particular type needed
2. Restrict object creation of some type to a single unit across entire program

## Uses
Including but not limited to
- Maintaining the same connection to a database to make each query
- Opening a SSH connection a server to do multiple tasks
- Limit access to a variable/space
- Limit the number of calls in some places

## Tests
- `TestGetInstance` checks that we actually get something when we ask for
  an instance of the counter
    + Creation of the object is delegated to some

