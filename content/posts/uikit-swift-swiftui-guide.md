---
title: "SwiftUI / UIKit / Swift Guide - Making iOS apps"
date: "2026-03-16"
published: true
---

## Swift

### Foundation

- `import Foundation` - core standard library for Apple platforms. Contains essential data types (Date, URL, Data)
- `UUID()` - "Universally unique identifier", 128-bit integer, represented by a specific string
- `Identifiable` - it is a protocol that a `struct` or `class` needs to follow and it says that an object must contain an `id`

```Swift
struct Habit : Identifiable {
    var id = UUID()
    var name: String
    var isCompleted: Bool = false
}
```

## Swift UI

### `View` protocol
- any UI element must conform to this protocol. Protocol defines rules
- The view protocol defines that a UI element must contain `body` property

### `import SwiftUI`
- to import Swift UI components, it replaces UIKit

### `some View`
- an "opaque return type"
- it basically means that this variable returns a `View` which is a must for a `struct` that conforms to the `View` protocol
- the `some` keyword just means that the View can be written in "simplified" language. For example `Text("Hi").padding()` does not return a `Text` object but a `ModifiedContent<Text, _PaddingLayout>` which can get very long and the `some` keyword abstracts this.

### `HStack`, `VStack`, `ZStack`
- arranges its childs, self explanatory

### `Image(systemName:)` SF symbols
- uses Apple's iconography to display icons

### `Spacer()`
- a flexible element that automatically expands to take the whole available remaining space.
- used for pushing elements to the top, bottom or aligning in `HStack` or `VStack`

### Declarative modifiers
- methods invoked on a view that modify its properties by returning a modified view

#### Sizing, layout

- `.padding()` - adds spacing
    - `.padding(.top, 16)`
- `.frame()` - defines dimensions of that specific view
    - `.frame(width:height:alignment)`
    - `.frame(maxWidth: .infinity)`

#### Styling

- `.shadow(color:radius:x:y:)` - adds drop shadow
- `.background()` - places another view behind this current view
- `.foregroundColor()` or `.foregroundStyle()` - changes color of Text or SF symbol in the view
- `.clipShape()` - masks the view to a specific range
    - `.cornerRadius(12)` can be replaced by `.clipShape(RoundedRectangle(cornerRadius: 12))`

#### Typography

- `.font()` - applies predefined typography, this makes sure that the apps responds correctly to accessibility
    - `.font(.title)`, `.font(.headline)`
- `.fontWeight()` - adjusts the thickness
    - `.fontWeight(.semibold)`

#### Interactivity

- `.onTapGesture { }` - executes code when this view is clicked

### State Management: The ViewModel
- to manage life cycle of the app and mutation of data
- 

## MVVM
- an architecture of an iOS called "Model View View Model"

## Sqlite-data

- It is a replacement of SwiftData, powered by SQL, built on top of GRDB
- [sqlite-data github](https://github.com/pointfreeco/sqlite-data)