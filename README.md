# flow as code

create a flow as code (in YAML format) is a utility that executes
workflows based on a yaml configuration you pass in.

Sample flow document: documentation/example-flow.yaml

_Remarks: the example document still communicates features not yet implemented_

## Usage

    flow-as-code --flow <flow.yaml> --name <test> [--debug]
    
    
    --flow  path to the file containing the flow definition
    --name  name of the step to execute (entrypoint)
    --debug Start printing debug messages

### Configuration

The configuration is done using a YAML document. It has the following structure:

    ---
    
    shell:       string      # binary in path or complete path to
                             # shell that will be called
    
    name:        string      # just give the thing a name
    description: string      # provide a more in depth description
    
    steps:       list        # provide a list of steps
    
A step contains the following properties:

    name:        string      # just give the thing a name
    on_success:  list        # follow up with steps on success (descendant)
    on failure:  list        # follow up with steps on failure (descendant)
    script:      list        # list of commands to execute

A descendant contains the following properties:

    name:        string      # must match the name of a step
    variables:   map         # A dictionary (map) of variables passed to
                             # the step to be executed

### Variable handling

The variables get special treatment as $VAR gets replaced by the environment
variable named VAR.

Also you can access the flow and the current step like this:

    {{.Flow.Name}}        # Name
    {{.Flow.Shell}}       # Shell
    {{.Flow.Description}} # Description
    {{.Flow.Steps}}       # All steps
    
    {{.Step.Name}}        # Name
    {{.Step.Script}}      # Commands to be executed
    {{.Step.OnSuccess}}   # List of descendants on success
    {{.Step.OnFailure}}   # List of descendants on failure

## History

|Version|Description|
|---|---|
|0.2.0|- support to pass variables to descendants|
||- license and contributor information|
||- integrated gops|
|0.1.0|initial minimal version|
