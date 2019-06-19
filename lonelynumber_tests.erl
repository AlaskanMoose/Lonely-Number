-module(lonelynumber_tests).

-include_lib("eunit/include/eunit.hrl").

find_test() ->
    true = 6 ==
	     lonelynumber:find([2, 3, 4, 1, 5, 3, 5, 2, 4, 6, 1]).
