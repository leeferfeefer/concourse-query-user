------------------------------------------------------------------------------------------------------------------------------ 
### Bridge Troll
------------------------------------------------------------------------------------------------------------------------------ 
     
        
### Want to use Concourse with input parameter functionality?
      
I wanted to make use of Concourse's dockerization ability to re-run failing pipeline steps without having to restart the entire pipeline process as Jenkins does. But, the downside of using a Concourse pipeline is no out of the box functionality that supports user input. So, I created Bridge Troll.
       
With Bridge Troll, one can leverage the benefits of a concourse pipeline while also having the same user input functionality supported by Jenkins. 
        
Bridge troll uses ssl connection to send credentials/data/whatever securely. 
      
       
        
     
      
        
------------------------------------------------------------------------------------------------------------------------------ 
### Bridge Troll in action:
------------------------------------------------------------------------------------------------------------------------------ 

![](https://github.com/leeferfeefer/concourse-query-user/blob/ssl_implementation/images/troll-concourse.png)

After going to https://localhost:3001:
![](https://github.com/leeferfeefer/concourse-query-user/blob/ssl_implementation/images/troll.png) 

After entering credentials:
![](https://github.com/leeferfeefer/concourse-query-user/blob/ssl_implementation/images/troll%20toll%20paid.png) 

Bridge Troll even has input validation:
![](https://github.com/leeferfeefer/concourse-query-user/blob/ssl_implementation/images/troll-validation.png) 

As you can see, the pipeline step will wait until a user has entered input before continuing.   
![](https://github.com/leeferfeefer/concourse-query-user/blob/ssl_implementation/images/troll%20concourse%20done.png)
       
        
        
------------------------------------------------------------------------------------------------------------------------------ 
### Guides
------------------------------------------------------------------------------------------------------------------------------ 
* [Follow this guide to create a Concourse darwin worker](https://github.com/leeferfeefer/concourse-query-user/wiki/Creating-a-darwin-worker)   
* [Follow this guide to create a self-signed certificate](https://github.com/leeferfeefer/concourse-query-user/wiki/Creating-a-self-signed-certificate)
    
    
     
------------------------------------------------------------------------------------------------------------------------------ 
#### Note:
------------------------------------------------------------------------------------------------------------------------------ 
* Using the example pipeline configuration assumes you have a darwin worker running
* In order for HTTPS communication to be established, a self signed certificate can be used. Simply generate one and have 
the UI and Server point to the certificate and key. 
