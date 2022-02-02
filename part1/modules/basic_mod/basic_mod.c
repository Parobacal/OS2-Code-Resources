#include <linux/module.h>
#include <linux/init.h>

MODULE_LICENSE(“GPL”);
MODULE_AUTHOR(“parobacal”);
MODULE_DESCRIPTION(“Basic Linux module.”);
MODULE_VERSION(“0.01”);

static int __init load_module(void) {
    printk(KERN_INFO "Hello world!\n");
    return 0;
}

static void __exit remove_module(void) {
    printk(KERN_INFO "Goodbye!\n");
}

module_init(load_module);
module_exit(remove_module);