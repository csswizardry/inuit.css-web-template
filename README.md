![inuit.css](http://inuitcss.com/img/content/logo.png)

# inuit.css web template

This is the inuit.css web template; the wrapper for the popular
[inuit.css](https://github.com/csswizardry/inuit.css) framework.

This template is a barebones and unopinionated directory structure which simply
allows you to include the inuit.css core library in your builds as an updatable
Git submodule.

To quickly install inuit.css, run the following commands:

    $ git clone --recursive git@github.com:csswizardry/inuit.css-web-template.git your-project-folder
    $ cd your-project-folder
    $ ./go

What we are doing here is cloning an instance of the inuit.css-web-template and
its submodules (that’s what the `--recursive` does) into a directory which you
specify. Next we `cd` into that directory and run [our `go` script](https://github.com/csswizardry/inuit.css-web-template/blob/master/go).
This script (courtesy of [Nick Payne](http://twitter.com/makeusabrew)) simply
removes the web template’s Git instance and replaces it with a fresh one for
your project, whilst also maintaining your inuit.css submodule.

For a more detailed overview on what inuit.css is, and how to install and use
it, please refer to the documentation in the README in
[the main inuit.css repository](https://github.com/csswizardry/inuit.css).
