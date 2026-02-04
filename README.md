# OS-HW0
Operating Systems HW0 
-------------------------------------------------------------------------------------------
Group Members: Katrina Wong, William Quinterogomez
-------------------------------------------------------------------------------------------
Part 1: Process and Inter-Process Communication (IPC)
- Instruction: This program creates two processes (producer and consumer)that communicates through pipe. The producer sends an number  and then the consumer will recieve the numbers and both prints numbers out in synchronized order. This program can be executed by running this file in the command line. 
- Design: The program is designed using the inter-process communication (IPC). Two processes, producer and consumer uses IPC to communicate with each other. The producer generates data in this case numbers, while the consumer receives and processes that data. In this case the program will execute 5 numbers in synchronized order from 1 to 5. The program terminates after 5 values have been both produced and consumed. 
