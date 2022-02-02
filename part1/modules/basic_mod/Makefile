obj-m += basic_mod.o

all:
	make -C /lib/modules/$(shell uname -r)/build M=$(PWD) modules

clean:
	make -C /lib/modules/$(shell uname -r)/build M=$(PWD) clean

test:
	sudo dmesg -C
 	sudo insmod basic_mod.ko
 	sudo rmmod basic_mod.ko
 	dmesg