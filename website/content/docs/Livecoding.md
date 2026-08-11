"Livecoding" means that when you talk to a Pipefish service in the REPL, if its source code has been changed since last you interacted with it, the service will automatically be re-initialized with the new code.

So if you also have autosave turned on in your editor (as we strongly recommend), then changes in your code will propagate through to the service without you needing to do anything. Try it out! If you're still running the service from the previous section, change `Hello` to `Hi there` in the script, and see what happens.

<img width="701" height="589" alt="image" src="https://github.com/user-attachments/assets/923c7177-ba34-4921-9105-775721fd3c79" />

Livecoding can be turned on and off with `hub live on` and `hub live off`.

Between the REPL and livecoding, it is very easy to just casually mess around with Pipefish and we recommend that you do so as you follow through this manual.