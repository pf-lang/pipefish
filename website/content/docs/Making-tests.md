This page lists all the hub commands you need to make and run tests.

#### `hub snap (scriptFilename string) as (testFilename string)`

This creates a test to let you record your interactions with the REPL: what you type and the service’s responses will be recorded. (There will be a `#snap →` prompt to remind you that this is what you’re doing.) 

While the test is underway, Pipefish will turn serialization on, so that you can tell the difference between "true" and true; and between four spaces and a tab, etc.

#### `hub snap (scriptFilename string)`

This supplies you with an unused default name for the test, if you don't supply one.

To finish recording, you tell Pipefish what to do with the snap, using one of the four following commands:

#### `hub snap good`

This means: Stop the test. This is the desired behavior and I want to make it into a test that will ensure it keeps happening.

#### `hub snap bad`

This means: Stop the test. his is undesirable behavior and I want to make a test that checks that it doesn’t happen again.

#### `hub snap record`

This means: Stop the test. I want to be able to replay this and see how the output changes as I change the script.

#### `hub snap discard`

This means: Stop the test. I don’t need this test after all, throw it away.

#### `hub test(scriptFilename string)`

This runs all the tests ended with `snap good` or `snap bad` and checks that they get the right results.

#### `hub replay(testFilename string)`

This replays the inputs from a test ended with `snap record`, and shows the new outputs.

#### `hub replay diff(testFilename string)`

This does the same as `hub replay` but shows the difference between the results obtained and those produced when the test was made.