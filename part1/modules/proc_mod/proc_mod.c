#include <linux/proc_fs.h>
#include <linux/seq_file.h>
#include <linux/module.h>
#include <linux/init.h>
#include <linux/mm.h>


MODULE_LICENSE("GPL");
MODULE_AUTHOR("parobacal");
MODULE_DESCRIPTION("Basic process information Linux module.");
MODULE_VERSION("0.01");

static int writeFile(struct seq_file* archivo, void *v){

    seq_printf(archivo, "==========================\n");
    seq_printf(archivo, "=           OS2          =\n");
    seq_printf(archivo, "=        parobacal       =\n");
    seq_printf(archivo, "=         proc_mod       =\n");
    seq_printf(archivo, "==========================\n\n");

    return 0;
}

static int atOpen(struct inode* inode, struct file* file){
    return single_open(file, writeFile, NULL);
}

static struct file_operations ops = {
    .open = atOpen,
    .read = seq_read
};

int proc_count(void)
{
  int i=0;
  struct task_struct *thechild;
  for_each_process(thechild){
    i++;
  }
  return i;
}

static int load_module(void) {
    printk(KERN_INFO "Total running processes: %d .\n", proc_count());

    proc_create("proc_mod", 0, NULL, &ops);
    return 0;
}

static void unload_module(void) {
    printk(KERN_INFO "Goodbye!\n");

    remove_proc_entry("proc_mod", NULL);
}

module_init(load_module);
module_exit(unload_module);