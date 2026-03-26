---
title: "SwiftUI / UIKit / Swift Guide - Making iOS apps"
date: "2026-03-16"
published: true
---

## Swift

### class

- it is a reference type, it ensures a single source of truth rather than a struct

### struct

- it is a value type
```Swift
struct Car : Identifiable {
    var id = UUID()
    var model: String
    var year: Int
}
```

### array

- `var cars: [Car] = []` - an empty array of objects

### `firstIndex(where:)`
- `cars.firstIndex(where: {$0.id == id}` - find the first car that matches a passed index

### `guard`
- `guard !newHabitName.IsEmpty else {return}` - making sure a field is not empty before proceeding

### Half-open range operator and map
- for loop
- maps stuff
```Swift
(0..<70).map {_ in Bool.random()}
```

### `Array(repeating:count:)`
- allocates a block of memory with identical elements

```Swift
let columns = Array(repeating: GridItem(.fixed(14), spacing: 4), count: 10)
```

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

### `Text`
- to display read-only text

### `Button`
- initiates action when clicked by user

### `Image`
- renders an image or an SF symbol

### `TextField`
- single-line string input

### `Toggle`
- boolean state switch

### `Section`
- grouping content within a List or a Form

### `Form`
- data-entry layouts

### `ScrollView`
- enables scrolling

### `TabView`
- switches between multiple child views - just a navigation bar
- `.tabItem { }` - defines label and image in the navigation bar.
- `Label(title:systemImage:)` - represents a user interface item in the TabView



```Swift
struct ContributionGridView: View {
    let columns = Array(repeating: GridItem(.fixed(14), spacing: 4), count: 10)
    
    var completionData: [Bool] = (0..<70).map {_ in Bool.random()}
    
    var body: some View {
        LazyHGrid(rows: Array(repeating: GridItem(.fixed(14), spacing: 4), count: 7), spacing: 4) {
            ForEach(0..<70, id: \.self) { index in
                RoundedRectangle(cornerRadius: 3)
                    .fill(completionData[index] ? Color.accentColor : Color.gray.opacity(0.2))
                    .frame(width: 14, height: 14)
            }
        }
        .padding()
        .background(Color(UIColor.secondarySystemGroupedBackground))
        .clipShape(RoundedRectangle(cornerRadius: 12, style: .continuous))
    }
}
```

### Grid item
`GridItem(.fixed(14), spacing: 4)`
- element of a grid

### LazyHGrid, LazyVStack
- grid or vstack but renders only when it is on screen

### `id: \.self`
- unique identifier to track each view dynamically

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
- `.opacity()` - from 0.0 to 1.0

#### Typography

- `.font()` - applies predefined typography, this makes sure that the apps responds correctly to accessibility
    - `.font(.title)`, `.font(.headline)`
- `.fontWeight()` - adjusts the thickness
    - `.fontWeight(.semibold)`

#### Interactivity

- `.onTapGesture { }` - executes code when this view is clicked

### View Model
- to manage life cycle of the app and mutation of data

```Swift
import SwiftUI
import Observation

@Observable
class HabitViewModel {
    var habits: [Habit] = []
    
    func addHabit(name: String) {
        let newHabit = Habit(name: name)
        habits.append(newHabit)
    }
    
    func toggleCompletion(for id: UUID) {
        if let index = habits.firstIndex(where: {$0.id == id}) {
            habits[index].isCompleted.toggle()
        }
    }
}
```

#### `import Observation`
- allows SwiftUI to track changes to data models

#### `@Observable` Macro
- Transforms a class into an observable object (at runtime!)
- It will make a SwiftUI view automatically rerender when properties from this class change, for example `habits` array
- replaces legacy `@Published` and `ObservableObject` wrappers

### Content View

- it is the root for the UI hierarchy
- it will instantiate the ViewModel and render other Views

```Swift
import SwiftUI

struct ContentView: View {
    @State private var viewModel = HabitViewModel()
    
    @State private var newHabitName: String = ""
    
    var body: some View {
        NavigationStack {
            VStack {
                HStack {
                    TextField("Enter new habit...", text: $newHabitName)
                        .textFieldStyle(.roundedBorder)
                    
                    Button("Add") {
                        guard !newHabitName.isEmpty else { return }
                        viewModel.addHabit(name: newHabitName)
                        newHabitName = ""
                    }
                    .buttonStyle(.borderedProminent)
                }
                .padding()
                
                List(viewModel.habits) { habit in
                    HabitRowView(habit: habit)
                        .onTapGesture {
                            viewModel.toggleCompletion(for: habit.id)
                        }
                        .listRowSeparator(.hidden)
                        .listRowBackground(Color.clear)
                }
                .listStyle(.plain)
            }
            .navigationTitle("HabitKeep")
        }
    }
}
```

#### `@State`
- a property wrapper that establishes the view as the owner of the data. It makes sure that data persists across rerenders.

```Swift
struct ContentView: View {
    @State private var viewModel = HabitViewModel()
    @State private var newHabitName: String = ""
    @State private var isSettingsPresented: Bool = false
    
    var body: some View {
        NavigationStack {
            ZStack {
                Color(UIColor.systemGroupedBackground)
                    .ignoresSafeArea()
                
                VStack {
                    HStack {
                        TextField("New habit...", text: $newHabitName)
                            .padding(12)
                            .background(Color(UIColor.secondarySystemGroupedBackground))
                            .clipShape(RoundedRectangle(cornerRadius: 12, style: .continuous))
                        
                        Button {
                            guard !newHabitName.isEmpty else {return}
                            viewModel.addHabit(name: newHabitName)
                            newHabitName = ""
                        } label: {
                            Image(systemName: "plus")
                                .font(.system(size: 16, weight: .bold))
                                .foregroundColor(.white)
                                .padding(12)
                                .background(Color.accentColor)
                                .clipShape(Circle())
                        }
                    }
                    .padding()
                    
                    ScrollView {
                        LazyVStack(spacing: 12) {
                            ForEach(viewModel.habits) { habit in
                                HabitRowView(habit: habit)
                                    .onTapGesture {
                                        viewModel.toggleCompletion(for: habit.id)
                                    }
                            }
                        }
                        .padding(.horizontal)
                        .padding(.bottom, 24)
                    }
                }
            }
            .navigationTitle("HabitKeep")
            .toolbar {
                ToolbarItem(placement: .topBarTrailing) {
                    Button {
                        isSettingsPresented = true
                    } label: {
                        Image(systemName: "gearshape.fill")
                            .foregroundColor(.primary)
                    }
                }
            }
            .sheet(isPresented: $isSettingsPresented) {
                SettingsView()
            }
        }
    }
}

#Preview {
    ContentView()
}
```

```Swift
import SwiftUI

struct SettingsView : View {
    @Environment(\.dismiss) private var dismiss
    
    var body : some View {
        NavigationStack {
            List {
                Section(header: Text("Preferences")) {
                    Toggle("Haptic Feedback", isOn: .constant(true))
                    Toggle("Notifications", isOn: .constant(false))
                }
                
                Section(header: Text("Premium")) {
                    HStack {
                        Text("HabitKeep Pro")
                        Spacer()
                        Text("Inactive")
                            .foregroundColor(.secondary)
                        
                    }
                }
            }
            .navigationTitle("Settings")
            .navigationBarTitleDisplayMode(.inline)
            .toolbar {
                ToolbarItem(placement: .topBarTrailing) {
                    Button("Done") {
                        dismiss()
                    }
                    .fontWeight(.semibold)
                }
            }
        }
    }
}

#Preview {
    SettingsView()
}
```

#### `@Environment(\.dismiss)`
- a property that reads the value from the environment and dismissing it removes it from the navigation stack.

#### `.sheet(isPresented: content)`
- presents a modal view when the bool is set to true

#### `NavigationStack`
- a container view that is standard to iOS navigation interface with animations, navigation bar and etc.
- it manages a stack of views

#### `Color(UIColor: .systemGroupedBackground)`
- a UIKit color, so the app adapts better to dark and light mode

#### `@main`
- This designates the entry point to application's execution

#### `#Preview`
- a macro that enables to render a view in an isolated canvas

```Swift
#Preview {
    ContentView()
}
```

### Creating better looking UI
- ditch the basic UI of iOS for more customizable experiene of Views

- Sroll View + Lazy VStack $\rightarrow$ better approach for custom list that List

```Swift
import SwiftUI

struct HabitRowView {
    var habit: Habit
    
    var body: some View {
        
    }
}
```

## MVVM
- an architecture of an iOS called "Model View View Model"

```md ios-app-architecture

```

## Sqlite-data

- It is a replacement of SwiftData, powered by SQL, built on top of GRDB
- [sqlite-data github](https://github.com/pointfreeco/sqlite-data)