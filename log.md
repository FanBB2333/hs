
## Sample Log
build task in 2 s 867 ms
Launching com.example.myapplication
$ hdc shell aa force-stop com.example.myapplication
$ hdc shell mkdir data/local/tmp/xxxx
$ hdc file send xxxx/xxx-signed.hap "data/local/tmp/xxxx" in 61 ms
$ hdc shell bm install -p data/local/tmp/xxxx  in 407 ms
$ hdc shell rm -rf data/local/tmp/xxxx
$ hdc shell aa start -a EntryAbility -b com.example.myapplication in 224 ms
Launch com.example.myapplication success in 1 s 299 ms