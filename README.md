golang version of mysql_escape_string()

refer to mysql C api's escape_string_for_mysql function : https://github.com/twitter/mysql/blob/865aae5f23e2091e1316ca0e6c6651d57f786c76/mysys/charset.c

forked from https://github.com/liule/golang_escape and add some test.

tested with some simple manual written data and compare with php5's mysql_escape_string.