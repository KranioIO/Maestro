# Symphony Yml Schema Reference

Reference document for symphony creation.

---

### File Struct Example

```yaml
symphony: here-is-the-name-of-your-symphony

trigger:
  type: scheduler
  cron: 0/2 * * * *

tasks:
  start-task:
    type: dummy

  middle-task-A:
    type: dummy
    depends: start-task

  middle-task-B:
    type: lambda
    depends: start-task
    payload:
      name: your-function-name
      data:
        id: dummy-data
        value: dummy-data

  middle-task-C:
    type: dummy
    depends: start-task

  end-task:
    type: dummy
    depends:
      - middle-task-A
      - middle-task-B
      - middle-task-C
```

#### `trigger`
---
Defines how a symphony is started, possible values:
- none (default) - *only triggered by hand*
- schedule - *triggered by a `cron` definition*
- event - *triggered by an async method (SQS, SNS)*
- symphony - *triggered when another symphony finish his execution without errors*

#### `tasks`
---
The group of tasks definitions following this structure:

```
task-name:
  task-field:
  task-field:

task-name:
  task-field:
  task-field:

...
```

### `task` fields

---

depends: *Accept a name or a list of dependency task names*

type:
  - dummy (default) - *only waits a random time to execute*
  - lambda - *invoke a lambda function synchronously*
    - name: *your function name for invokation*
    - data (nullable): *the payload sent to your function* (parsed to json at runtime)
