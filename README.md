# concourse-query-user

### Want to use Concourse with input parameter functionality?

I wanted to make use of Concourse's dockerization ability to re-run failing pipeline steps without having to restart the entire pipeline process as Jenkins does. But, the downside of using a Concourse pipeline is no out of the box functionality that supports user input. With this example, one can leverage the benefits of a concourse pipeline while also having the same user input functionality supported by Jenkins.   


  
    
User query in action:

![](https://github.com/leeferfeefer/concourse-query-user/blob/osascript/images/User%201%20in%20progress.png)
![](https://github.com/leeferfeefer/concourse-query-user/blob/osascript/images/User%202%20in%20progress.png) 

As you can see, the pipeline step will wait until a user has entered input before continuing.   

![](https://github.com/leeferfeefer/concourse-query-user/blob/osascript/images/complete.png)




[Follow this guide to create a Concourse darwin worker](https://github.com/leeferfeefer/concourse-query-user/wiki/Creating-a-darwin-worker) 


#### Note:
* This solution makes use of osascript which only works for Mac:
  * A solution for Windows has not been worked on yet
* Using the example pipeline configuration assumes you have a darwin worker running
