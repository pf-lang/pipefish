For a standard setup, use the following three commands in order. (But omit the first if you've already set up your SQL interop.)

#### `hub config db`

This will ask you about the SQL server the hub should use as a database. If you don't have one, you'll need to set one up before going further.

#### `hub config admin`

This will ask you for your details (username, password, email, etc) so that it can register you as the first admin of the hub and grand high ruler of the services.

#### `hub listen(path string, port int)`

Having completed the previous step you could leave it as it is and you'd have an administered server running on your desktop. It could certainly serve multiple users, but only if they take it in turns to log on, because they'd have to share the one keyboard.

`hub listen` supplies the final step by telling the hub which path and port to get its information from. At this point anyone who wants to talk to it, including you, will have to use a client such as HubTalk. (Or, strictly speaking, you could use `curl` and roll your own JSON, but life is short and you have other things to do.)