
#!/bin/bash

ytt -f simple-demo -f ../../values.yml | kbld -f - | kapp deploy -a simple-demo -f- -c -y
