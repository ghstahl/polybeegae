# Atlanta
[TheWolf]: /images/wolf.jpg "Optional title attribute"

<div class="wolf">
![Flowers][TheWolf] 
</div>


### Get Started

For the "A tutorial is all I need crowd": **[Enough chit chat, Lets do this thing!](docs/tutorial/jump-start/)**.  

### Motivation

I rarely need to create a console app, but it is a requirement to know how to properly create one to be consumed by clients.  In my case, these are usually Developers, Build Engineers, QA Engineers, and Operations.  After a few of these, I got tired to writing a help section, and trying to yet again pick the right open source parsers.  When finally looking at the the work that went into the app, most of it wasn't for the reason that I needed to build the app.  A console app should be something before you actually put in your reason for wanting one, i.e. your code to transform a json file into xml is harldy anything when compared to the help section.


1. This is the last time I am writing a help display section for an app!
2. This is the last time I am picking a parser!

### Opinions

1. Parsing commands line arguments has nothing to do with help.  
    Parsers should parse and not pollute themselves with adding descriptions strings to commands.  My favorite parser, [Pingo Variant of Fluent Command Line Parser!](https://github.com/ghstahl/fluent-command-line-parser), is simple and even simpler when you don't engage in its help sub system.  If I had my way, I would completely remove any notion of help from that project.

2. Help is just another command.  
    Just like any other command, it needs its own parser.  The [Pingo.CommandLineHelp Project](https://github.com/ghstahl/PingoConsole/tree/master/Pingo.CommandLineHelp) implements the [ICommand interface](https://github.com/ghstahl/PingoConsole/blob/master/Pingo.CommandLine/Contracts/Command/ICommand.cs) and [ICommandHelp interface](https://github.com/ghstahl/PingoConsole/blob/master/Pingo.CommandLine/Contracts/Help/ICommandHelp.cs) just like every other command implementation.
    
3. If you use this, you will only be asked to write code that nobody has written before.
    All you need to write is your command and provide the framework your help strings.  
    
4. I should be able to make the mother of all command line apps by simply installing a bunch of nugets that implement the Pingo Console command interfaces.
    [Create Your Own Command!](docs/tutorial/create-command-plugin/).
    
5. Command-Line apps should be single stand-alone executables.  


### References

The help section in Pingo Console is exactly like nuget.exe.  So much so, that if Pingo Console was around I would hope the nuget.exe folks would have used it to jump stater their project.
