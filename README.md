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

## History

|Version|Description|
|---|---|
|0.1.0|initial minimal version|
