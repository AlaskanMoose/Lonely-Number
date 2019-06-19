-module(lonelynumber).

-compile(export_all).

find(List) ->
    lists:sum(sets:to_list(sets:from_list(List))) * 2 -
      lists:sum(List).
