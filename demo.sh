#!/bin/bash

cd ~/chia-blockchain
. ./activate

cd

# for more than one temp dir
#tp_num=1
n=1
for final in a b c d e f g h a b c d e f g h a b c d e f g h
do
  nohup chia plots create -t /mnt/tp3 -p 8c298823361cd86ae19a02389dc991744a351908e39045594894d3a027da2c5a55d1bcfe26fdb8d2da9cc70533a5a981 -f 94812e8f283287e7226bc396d292d2a68af31eb64c0cc3cf679de562e351f666c14b9f3c478722bacf3f80acad08811a -n 20 -u 64 -b 6800 -k 32 -r 2 -d /mnt/data$final >> create$n.log &
  # for more than one temp dir
#  let rem=$(($n%3))
#  if [ $rem = 0 ]; then
#      let tp_num=$tp_num+1
#  fi
  let n=$n+1
done


cp1=0
cp2=1
for p in $(ps -ef | grep create | grep -v grep | awk '{print $2}')
do
#  let cp2=$cp1+24
  taskset -cp $cp1,$cp2 $p
  let cp1=$cp1+2
  let cp2=$cp2+2
done


# 194
# 24 24

#!/bin/bash

. ./chia-blockchain/activate

n=1
for final in a b c d e f g h a b c d
do
  nohup chia plots create -t /mnt/tp2 -p 8c298823361cd86ae19a02389dc991744a351908e39045594894d3a027da2c5a55d1bcfe26fdb8d2da9cc70533a5a981 -f 94812e8f283287e7226bc396d292d2a68af31eb64c0cc3cf679de562e351f666c14b9f3c478722bacf3f80acad08811a -n 20 -u 32 -b 20400 -k 33 -r 4 -d /mnt/middle_tmp >> create$n.log &
  let n=$n+1
done


cp1=0
cp2=1
for p in $(ps -ef | grep create | grep -v grep | awk '{print $2}')
do
  let cp3=$cp1+24
  let cp4=$cp2+24
  taskset -cp $cp1,$cp2,$cp3,$cp4 $p
  let cp1=$cp1+2
  let cp2=$cp2+2
done




#!/bin/bash

. ./chia-blockchain/activate

tp_num=1
n=1
for final in a b c d e a b c d e a b
do
  nohup chia plots create -t /mnt/tp$tp_num -p 8c298823361cd86ae19a02389dc991744a351908e39045594894d3a027da2c5a55d1bcfe26fdb8d2da9cc70533a5a981 -f 94812e8f283287e7226bc396d292d2a68af31eb64c0cc3cf679de562e351f666c14b9f3c478722bacf3f80acad08811a -n 20 -u 32 -b 20000 -k 32 -r 4 -d /mnt/data$final >> create$n.log &
  let rem=$(($n%3))
  if [ $rem = 0 ]; then
      let tp_num=$tp_num+1
  fi
  let n=$n+1
done

cp1=0
cp2=1
for p in $(ps -ef | grep create | grep -v grep | awk '{print $2}')
do
  let cp3=$cp1+24
  let cp4=$cp2+24
  taskset -cp $cp1,$cp2,$cp3,$cp4 $p
  let cp1=$cp1+2
  let cp2=$cp2+2
done





#!/bin/bash

 . ./chia-blockchain/activate


tp_num=1
n=1
for final in e f g h i e f g h i e f g h i e f g h i e f g h
do
  nohup chia plots create -t /mnt/tp$tp_num -p 8c298823361cd86ae19a02389dc991744a351908e39045594894d3a027da2c5a55d1bcfe26fdb8d2da9cc70533a5a981 -f 94812e8f283287e7226bc396d292d2a68af31eb64c0cc3cf679de562e351f666c14b9f3c478722bacf3f80acad08811a -n 20 -u 64 -b 10000 -k 32 -r 2 -d /mnt/data$final >> create$n.log &
  let rem=$(($n%3))
  if [ $rem = 0 ]; then
      let tp_num=$tp_num+1
  fi
  let n=$n+1
done


cp1=0
#cp2=1
for p in $(ps -ef | grep create | grep -v grep | awk '{print $2}')
do
  let cp2=$cp1+24
  taskset -cp $cp1,$cp2 $p
  let cp1=$cp1+1
#  let cp2=$cp2+2
done



#!/bin/bash
while true
do
  n=$(ps -ef | grep chia-plotter | grep -v grep |wc -l)
  if [ "$n" -lt 5 ];then
     nohup .~/chia-plotter/build/chia_plot 8c298823361cd86ae19a02389dc991744a351908e39045594894d3a027da2c5a55d1bcfe26fdb8d2da9cc70533a5a981 94812e8f283287e7226bc396d292d2a68af31eb64c0cc3cf679de562e351f666c14b9f3c478722bacf3f80acad08811a /mnt/tp1/ /mnt/tp1/  48 5 >> log$n.log &
  fi
  sleep 60
done
