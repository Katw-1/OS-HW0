# OS HW 0: Part 1 
#Group Members: Katrina Wong, William Quinterogomez
# Instruction: This program creates two processes (producer and consumer)that communicates through pipe. The producer sends an number 
#  and then the consumer will recieve the numbers and both prints numbers out in synchronized order. This program can be executed by 
#  running this file in the command line. 

from multiprocessing import Process, Pipe #importing inter-process communication 
import time #importing time so there will be a pause 

def producer(conn): # defining producer, conn is from pipe
    for i in range(1, 6):   # range is from 1 to 0 since it starts at 0
        print(f"Producer: {i}")  #prints Producer: and any i number in the for loop in range of (1,6) 
        conn.send(i)        # sends the number to consumer
        time.sleep(1)     # a delay to make sure the number is sending to the consumer
    conn.close()

def consumer(conn): #defining consumer
    for num in range(5):  
        num = conn.recv()   # defining var num and blocking recieve
        print(f"Consumer: {num}")  #prints Consumer: with the number recieved from the producer
    conn.close()

if __name__ == "__main__": #creates pipe for inter-communication process
    parent, child = Pipe() #parent is used for producer and child is used for consumer

    p = Process(target=producer, args=(parent,)) #creates and executes producer functions
    c = Process(target=consumer, args=(child,)) #creates and executes comsumer functions

    p.start() #starts the producer
    c.start() #starts the consumer

    p.join() #this waits for the producer process
    c.join() #this waits for the consumer process