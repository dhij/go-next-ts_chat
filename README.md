# Realtime Chat

This is a realtime chat built with Go, Next, and Typescript.

The tutorial is available [here](https://www.youtube.com/watch?v=W9SuX9c40s8)

## Hub Architecture

![Initial Hub Architecture](/assets/hub_initial.jpg)

First, we have the hub running on a separate goroutine which is the central place that manages different channels and contains a map of rooms. The hub has a Register and an Unregister channel to register/unregister clients, and a Broadcast channel that receives a message and broadcasts it out to all the other clients in the same room.

![Client joins room](/assets/join_room.jpg)

A room is initially empty. Only when a client hits the `/ws/joinRoom` endpoint, that will create a new client object in the room and it will be registered through the hub's Register channel.

![Hub Architecture](/assets/hub_architecture.jpg)

Each client has a `writeMessage` and a `readMessage` method. `readMessage` reads the message through the client's websocket connection and send the message to the Broadcast channel in the hub, which will then broadcast the message out to every client in the same room. The `writeMessage` method in each of those clients will write the message to its websocket connection, which will be handled on the frontend side to display the messages accordingly.
