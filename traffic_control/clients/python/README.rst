************
Introduction
************
This is the Traffic Ops Python Client for Python 2.x and Python 3.x.

.. attention:: Traffic Control version 3.0.0 officially deprecates Python 2.x support. Starting with Traffic Control version 4.0.0, Python 2.x support will be dropped, and only Python 3.x will be supported. Users and developers are encouraged to switch to Python 3 as soon as possible.

.. note:: This client has only been tested against Python 2.7 and Python 3.6. Other versions may or may not work.

Installation
============
The official installation method is to use ``pip`` to install directly from from GitHub.

.. code-block:: shell
	:caption: Install Using 'pip' From GitHub

	pip install git+https://github.com/apache/trafficcontrol.git#"egg=trafficops&subdirectory=traffic_control/clients/python/trafficops"

	# or
	# pip install git+ssh://git@github.com/apache/trafficcontrol.git#"egg=trafficops&subdirectory=traffic_control/clients/python/trafficops"

Local Installation
------------------
The preferred method is to use ``pip`` to install locally. Starting from the repository's root directory, the following script should do the job:

.. code-block:: shell
	:linenos:
	:caption: Local Installation of Python Client

	cd traffic_control/clients/python
	pip install .

	# The above will install using the system's default Python interpreter - to use a specific
	# version it will be necessary to specify the interpreter and pass the 'pip' module to it.
	# e.g. for the system's default Python 2 interpreter, typically one would do:
	# sudo -H /usr/bin/env python2 -m pip install .

	# Developers may wish to use the '-e' flag. This will install the package 'edit-ably',
	# meaning that changes made to the package within the repository structure will be effected on
	# the system-wide installation.
	# sudo -H pip install -e .

	# If your system does not have 'pip', but does (somehow) have 'setuptools' as well
	# as the package's dependencies, you can call 'setup.py' directly to install the
	# package for the system's default Python 3 interpreter
	# sudo -H ./setup.py install

The local installation method requires ``pip`` and ``setuptools``. ``setuptools`` should be installed if your system has ``pip``, but if you are missing either of them they can both be relatively easily installed with Python standard libraries.

.. code-block:: shell
	:caption: Setuptools Package Installation

	# Here I'm using 'python' because that points to a Python 3 interpreter on my system. You may
	# wish to use 'python3' or 'python2' (please not that one) instead.
	sudo -H python -m ensure_pip
	sudo -H python -m pip install -U pip

	# If your system's 'python' already has 'pip', then you may skip to this step to
	# install only 'setuptools'
	sudo -H python -m pip install setuptools

Development Dependencies
------------------------
To install the development dependencies, first ensure that your system has ``pip`` and ``setuptools`` then use ``pip`` to install the development environment.

.. code-block:: shell
	:caption: Development Dependencies Installation

	pip install -e .[dev]
