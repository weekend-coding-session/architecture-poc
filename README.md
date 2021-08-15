# A PoC of clean Go project architecture 
This repo is meant to build a PoC of Go project architecture that follows the concept of clean architecture by Uncle Bob.

It should contain most of the basic functions needed in a production web server and we can always follow this project architecture in the future.

## Requirements/ Specs and Progress
- [x] **IoC** - It should follow Inversion of Control (each layer is dependent on Abstraction)
- [x] **DI** - Each layer is using dependency injection
- [ ] **Routers&Controllers** - It should have a cleaner router abstractions. Registering controller should be easier, the endpoint route should be defined together with the controller function. (I should see the route when I am looking at my controller function)
- [ ] **Endpoint declaration** - For routes definition, it should have "controller routes" and "handler routes" (an endpoint = controller route + handler route)
- [ ] **HTTP delivery** - Http related stuffs, constants should be abstracted with a cleaner interface. With current gorilla mux, getting input params, sending outputs with http metadata etc. are tedious
- [ ] **Error constants** - A better Error abstraction, with error codes, error message etc.
- [ ] **Error handling** - Error handling on http layer should be handle by error handling middleware (Self written), possibly a serveNext middleware. (It will work like node.js's next())
- [ ] **Logging** - A structured logger. Uber's zap logging should be good
- [ ] **gRPC** - It should have another gRPC layer (server and client)
- [ ] **Data model** - Should have a clear definition of how data model is used, where to define, when to use which etc. (there are DTOs, DB models, business entities, common types etc.)
- [ ] **Mocks** - Mock objects should be easily created for any layers. Have a script that will automatically update the mock objects.
- [ ] **Server init** - A more elegant way to instantiate all the controller, service, db instance (Probably a better function to wrap everything will do. Or maybe store all the instance under a big struct)