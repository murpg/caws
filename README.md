# CAWS

#
## CommandBox as Windows Service

CommandBox is one of the most important contributions to the ColdFusion community since the introduction of ColdFusion MX. What is CommandBox? CommandBox is a standalone native development tool for all 3 operating systems. It provides you with a CLI for developing all types and classes of applications. It also has the capability for you to choose ala carte an embedded server that is capable of running Lucee or Adobe ColdFusion. It seamlessly integrates with any of the other Ortus products.  It is also very flexible which allows it to also be integrated with many other products as you will see. CommandBox needs the Java in some form to run. This could be installed natively or in the same folder as the box.exe file. There is also a need to be able to install CommandBox as a Windows Service. What are the benefits of this.

1. Install once and never worry about installing again.
2. Securely share an application to none tech folks that only have a windows environment.
3. Update the application without none tech intervention.
4. Easily have new releases.

#
## Resources Needed

1. [Download CommandBox with JRE](https://www.ortussolutions.com/products/commandbox)
Put this in some folder on a thumb drive for this example I will use I: drive then inside of this folder you want to do 3 things. Create a file called commandbox.properties inside this file place this text. It is going to me your CommandBox home directory. Here is the text. commandbox\_home=I:\\ODWCAWS\\home\\.CommandBox
double back slashes are needed because it needs to be escaped.
2. [Download the Go language](https://golang.org/) this is only if you want to recompile the caws.exe file. Here are instructions for that here.
3.
3.Create 2 empty folders called home and websites. The home folder is where the variable commandbox\_home variable is pointing to in the commandbox.properties file.
Once done your folder structure should look like this.
 ![directory](https://user-images.githubusercontent.com/530964/32448028-c7008990-c2db-11e7-95c7-dd9d1922d167.jpg)
4. In the website directory pull down the repository from
 [https://github.com/murpg/caws](https://github.com/murpg/caws)
Create a folder to hold your website. It can be named anything. Add this folder to your .gitignore file and update your server.json
at ODWCAWS\websites\server.json  
&quot;openBrowserURL&quot;:&quot;/whateveryounamedfolder&quot;
5. [Download NSSM](https://nssm.cc/download) to setup a Windows Service create a folder called utilities at [\\ODWCAWS\utilities](./../../%5C%5CODWCAWS%5Cutilities)
Follow the steps below  
 ![](data:image/*gwUEIAABCECgqwT++X2/8v8BgbFhL1DoikQAAAAASUVORK5CYII=)
![nssm1](https://user-images.githubusercontent.com/530964/32448281-82657c04-c2dc-11e7-85a9-f0f1af1ebb92.jpg)
![nssm2](https://user-images.githubusercontent.com/530964/32448304-924c8db0-c2dc-11e7-97ab-bf6a930fee20.jpg)
![nssm3](https://user-images.githubusercontent.com/530964/32448346-a8ec8250-c2dc-11e7-992a-b942fee3d717.jpg)
Important that you add an account with admin privileges  
![nssm4](https://user-images.githubusercontent.com/530964/32448577-460c8b70-c2dd-11e7-9c7d-4677b91686e4.jpg)
![nssm5](https://user-images.githubusercontent.com/530964/32448598-557962d6-c2dd-11e7-8331-fa134fde72af.jpg)  
Click the okay button. If you have followed all steps you will see a button that says CAWS successfully started.
6. Next CD to the directory where you installed caws.exe run caws start if on CMD run .\caws start if on Powershell