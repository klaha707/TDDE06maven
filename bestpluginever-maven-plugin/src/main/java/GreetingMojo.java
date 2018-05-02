package main.java;

import org.apache.maven.plugin.AbstractMojo;
import org.apache.maven.plugin.MojoExecutionException;
import org.apache.maven.plugins.annotations.Mojo;
import org.apache.maven.plugins.annotations.Parameter;

import java.io.*;
 
/**
 * Says "Hi" to the user.
 *
 */


@Mojo( name = "sayhi")
public class GreetingMojo extends AbstractMojo
{
	@Parameter( property = "sayhi.filename", defaultValue = "file.txt" )
	private String filename;
    public void execute() throws MojoExecutionException
    {
        getLog().info( " ##### Hello, world! ##### " );
        String text = "Hello TDDE06";
        BufferedWriter output = null;
        try {
        	File myFile = new File(filename);
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