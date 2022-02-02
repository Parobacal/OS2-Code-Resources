#include <linux/module.h>
#include <linux/init.h>

static int __init load_module(void) {
    printk(KERN_INFO "Hello world!\n");
    return 0;
}

static void __exit remove_module(void) {
    printk(KERN_INFO "Goodbye!\n");
}

module_init(load_module);
module_exit(unload_module);