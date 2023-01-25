# github-playground

## Workflow

### another repositroy workflow invoker

invoke one repository's GitHub Actions from another repositry.

1. 以下の権限を持った GitHub App を作成する
  - Read and write access to actions
  - Read access to metadata

2. 作成した GitHub App を呼び出される側のリポジトリに install する
3. 呼び出される workflow dispatch workflow を作成する

GitHub Actions によって上記の workflow を呼び出す場合

1. 呼び出し側のリポジトリに上記の GitHub App をインストールする（トークン取得を自作する場合、不要かもしれない）
2. 呼び出される側の workflow を呼び出す workflow を作成する
  - [expample here](./.github/workflows/invoke.yml)

## Mermaid

playground

```mermaid
graph TD
  A --> B;
  B --> C;
  A --> C;
  C --> D;
  D --> B;
```

```mermaid
sequenceDiagram
  participant A
  participant B
  
  A --> B: do
  loop check
    B --> B: self check
  end
  B --> C: notify
  B --> A: done
```

```mermaid
gantt
  title Sample
  section Section1
  run1: a1, 2022-02-17, 30d
  run2: a2, 2022-02-15, 1d
  run3: a3, 2022-02-14, 3d
  run4: a4, 2022-02-18, 2d
  
  section Section2
  do1: b1, 2022-03-17, 2d
  do2: b2, 2022-03-15, 1d
  do3: b3, 2022-03-14, 10d
  do4: b4, 2022-03-18, 3d
```

```mermaid
stateDiagram
  [*] --> active
  active --> sleep
  active --> run
  run --> active
  sleep --> active
  sleep --> [*]
```

## Test
