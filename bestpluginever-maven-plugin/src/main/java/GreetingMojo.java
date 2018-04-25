package main.java;

import org.apache.maven.plugin.AbstractMojo;
import org.apache.maven.plugin.MojoExecutionException;
import org.apache.maven.plugins.annotations.Mojo;
import java.io.*;
 
/**
 * Says "Hi" to the user.
 *
 */


@Mojo( name = "sayhi")
public class GreetingMojo extends AbstractMojo
{
    public void execute() throws MojoExecutionException
    {
        getLog().info( "Hello, world." );
        /**
         * My File.
         */
        
        String text = "Hello TDDE06";
        BufferedWriter output = null;
        try {
        	File myFile = new File("file.txt");
        	output = new BufferedWriter(new FileWriter(myFile));
        	output.write(text);
        } catch (IOException e) {
        	e.printStackTrace();
        } finally {
        	if(output != null) {
        		try {
        			output.close();
        		} catch (IOException e) {
                	e.printStackTrace();
                }
        	}
        }
    }
}