Entangle examples for Go
========================

This repository contains examples of using Entangle services in Go.


Building
--------

To build the examples, make sure that you have `Entangle <https://github.com/entangle/entangle>`_ installed. Building should then be a simple case of cloning the repository and using ``make``:

::

   $ git clone --recursive git@github.com:entangle/example-go.git
   $ cd example-go
   $ make


Examples
--------

A simple arithmetic service
~~~~~~~~~~~~~~~~~~~~~~~~~~~

This example is a simple arithmetic service which provides integer arithmetic. The ``bin/arithmetic-server`` executable runs a server on port 5555 while the ``bin/arithmetic-client`` executable connects to the server and runs examples of using the arithmetic server.
