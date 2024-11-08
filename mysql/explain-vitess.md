# Use `vexplain all` instead of `EXPLAIN format=vtexplain`

Using `vexplain all <query>;` give more usefull output in one step over `EXPLAIN format=vtexplain <query>;`. 

It gives information about shard lookups AND the mysql explain output of the shard.  
