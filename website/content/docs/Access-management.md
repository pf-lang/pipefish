Besides being able to start and stop services, the administrators need to be able to add and remove users from groups, allow groups to use services, etc.

#### `create (grp string)`

Creates a new user group.

#### `uncreate (grp string)`

Deletes it. (For consistency, all the hub verbs are negated with an `un` prefix.)

#### `let(grp string) use (srv string)`

Lets a group use a service. There is no need for the service to be running when you do this: if a user in the group lists the services available to them an uninitialized service will show up as inactive.

#### `unlet(grp string) use (srv string)`

Makes a group no longer able to use a given service.

#### `add(usr string) to (grp string)`

Adds the user to the group.

#### `unadd(usr string) to (grp string)`

Removes the user from the group

#### `let(usr string) own (grp string)`

Makes a user into a group owner. This allows them to add and remove users from the group.

#### `unlet(usr string) own (grp string)`

Makes the user no longer an owner of the given group.

#### `groups of service(srv string)`

Lists all the groups that belong to a given service.

#### `groups of user(usr string)`

Lists all the groups that a given member belongs to.

#### `users of group(grp string)`

Lists all the users that belong to a given group.

#### `users of service(srv string)`

Lists all the users that have access to a given service.

#### `services of group(grp string)`

Lists all the services available to a given group.

#### `services of user(usr string)`

Lists all the services available to a given user.
