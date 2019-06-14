# concourse-query-user

### Want to use Concourse with input parameter functionality?

By using this example as a guide for your own pipeline configuration, there is no need to depend on Jenkins to finish the 
final production steps. For my team, we wanted to make use of Concourse's dockerization ability to re-run failing pipeline
steps without having to restart the entire pipeline process as Jenkins does. But, the downside of using a Concourse pipeline
is no out of the box functionality that supports user input. For this reason, we chose to use Jenkins to take care of the
process of pushing code to production with the use of developer credentials gathered from user input. The downside 
to this set-up is having a production pipeline split between ci-cd providers. This not only leads to managing code in 2 separate 
areas, but also code written in 2 different formats. Now, with this example, a team can leverage the benefits of a concourse
pipeline while also having the same user input functionality supported by Jenkins.   


  
    
User query in action:

![](https://github.com/leeferfeefer/concourse-query-user/blob/master/images/User%201%20in%20progress.png)
![](https://github.com/leeferfeefer/concourse-query-user/blob/master/images/User%202%20in%20progress.png) 

As you can see, the pipeline step will wait until a user has entered input before continuing.   

![](https://github.com/leeferfeefer/concourse-query-user/blob/master/images/complete.png)




Follow this guide to create a Concourse darwin worker: https://github.homedepot.com/dxf2ci5/concourse-query-user/wiki/Creating-a-Concourse-Darwin-Worker


#### Note:
* This solution makes use of osascript which only works for Mac:
  * A solution for Windows has not been worked on yet
* Using the example pipeline configuration assumes you have a darwin worker running

