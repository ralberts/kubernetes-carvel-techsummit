# Simple App

## Ytt

ytt -f simple-demo -f ../../values.yml

## Kbld

kbld -f simple-demo

## Kapp

kapp deploy -a simple-demo -f- -c -y

## Chained

ytt -f simple-demo -f ../../values.yml | kbld -f - | kapp deploy -a simple-demo -f- -c -y
