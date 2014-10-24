SSLTest
=======

SSLTest is a command line implementation of the fantastically useful [SSLTest](https://www.ssllabs.com/ssltest) from Qualys SSL Labs. Motiviation behind this implimentation is to allow a similar test suite to be run against internal SSL endpoints.


Example
-------

	$ ssltest --hostname gocardless.com 2>&-
	========== SSL Report ==========
	Hostname: gocardless.com:443

	TLS Versions
	------------
	SSLv3 : false
	TLS1.0: true
	TLS1.1: true
	TLS1.2: true

	TLS Ciphers
	-----------
	TLS_RSA_WITH_AES_256_CBC_SHA              : true
	TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA      : false
	TLS_ECDHE_RSA_WITH_RC4_128_SHA            : false
	TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA        : true
	TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256     : true
	TLS_RSA_WITH_RC4_128_SHA                  : false
	TLS_RSA_WITH_AES_128_CBC_SHA              : true
	TLS_ECDHE_ECDSA_WITH_RC4_128_SHA          : false
	TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA      : false
	TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA       : true
	TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA        : true
	TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256   : false
	TLS_RSA_WITH_3DES_EDE_CBC_SHA             : true

	Started At : 2014/10/24 13:24:03 BST
	Finished At: 2014/10/24 13:24:04 BST
	================================


Usage
-----

	--hostname <localhost> hostname/IP Address to test against
	--port     <443>       port on host
	--json                 output as JSON rather than human-readable report


TODO
----

 - run tasks concurrently (with --worker flag)
 - include certificate peer chain information
 - add non-golang default ciphers
 - simulate handshakes of multiple platforms
 - scoring (like Qualys SSLTest)


License
-------

	Copyright (c) 2014, James Cunningham
	All rights reserved.

	Redistribution and use in source and binary forms, with or without
	modification, are permitted provided that the following conditions are met:
	    * Redistributions of source code must retain the above copyright
	      notice, this list of conditions and the following disclaimer.
	    * Redistributions in binary form must reproduce the above copyright
	      notice, this list of conditions and the following disclaimer in the
	      documentation and/or other materials provided with the distribution.
	    * Neither the name of the <organization> nor the
	      names of its contributors may be used to endorse or promote products
	      derived from this software without specific prior written permission.

	THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
	ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
	WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
	DISCLAIMED. IN NO EVENT SHALL <COPYRIGHT HOLDER> BE LIABLE FOR ANY
	DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
	(INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
	LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
	ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
	(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
	SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

