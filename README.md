# OS-HW0
Operating Systems HW0 
-------------------------------------------------------------------------------------------
Group Members: Katrina Wong, William Quinterogomez
-------------------------------------------------------------------------------------------
Part 1: Process and Inter-Process Communication (IPC)
- Instruction: This program creates two processes (producer and consumer)that communicates through pipe. The producer sends an number  and then the consumer will recieve the numbers and both prints numbers out in synchronized order. This program can be executed by running this file in the command line. 
- Design: The program is designed using the inter-process communication (IPC). Two processes, producer and consumer uses IPC to communicate with each other. The producer generates data in this case numbers, while the consumer receives and processes that data. In this case the program will execute 5 numbers in synchronized order from 1 to 5. The program terminates after 5 values have been both produced and consumed. 

Part 2 Design: 
-I created a struct called Stack. In defining the Stack, I created an integer slice called data. Two method are created to add and remove elements from the stack, Push and Pop respectively. Push method appends the value that was in the parameter in to the Stack that was being ponted. Pop() method has more to it. The Stack that is being referred to by the method simply removes the last element by reassigning the numbers of elements in the stack exculding the last element, thus removing the final element.In the main function, I do two four loops that iterate 100 times. The first one isadding values starting from 0 to 100 into the stack while simultaneously display the value that was inputted to the Stack.The second one uses Pop() to remove elements one by one until stack is empty.
