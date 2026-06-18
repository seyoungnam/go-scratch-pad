# Day 1: Go Interfaces & Polymorphism

In Go, interfaces are **implicit**. Unlike Java or C++, there is no `implements` keyword. If a type defines all the methods declared by an interface, it implements that interface automatically.

---

## 💡 Key Conceptual Areas for Interviews

### 1. Pointer vs. Value Receivers
*   **Value Receiver `(t T) Method()`:** The interface can be satisfied by both value types `T` and pointer types `*T`. Go can automatically dereference a pointer to call the value method.
*   **Pointer Receiver `(t *T) Method()`:** The interface is **only** satisfied by pointer types `*T`. A value of type `T` cannot satisfy the interface because the method may modify the receiver, and copy-by-value would discard these changes. If the value is non-addressable, Go cannot take its address.

### 2. Under the Hood: Internal Structure
An interface is represented by a two-word pair under the hood (in `runtime` package):
*   **`eface` (empty interface `any` or `interface{}`):**
    *   Word 1: Pointer to type information (`_type`).
    *   Word 2: Pointer to the actual data (`data`).
*   **`iface` (non-empty interface with methods):**
    *   Word 1: Pointer to an interface table (`itab`). The `itab` holds the interface's static type, the concrete type's type information, and a list of function pointers for the interface's methods.
    *   Word 2: Pointer to the actual data (`data`).

### 3. The Dangerous Gotcha: "Typed Nil" Interface
An interface is only `nil` when **both** its type (`itab` / `_type`) and its value (`data`) are `nil`.
If you assign a concrete `nil` pointer to an interface variable, the interface variable itself **is NOT nil**!

```go
var mp *Megapack = nil
var dev Device = mp

if dev == nil {
    // This block is NOT executed!
    // dev is not nil because its type is *Megapack.
}
// Calling dev.FetchTelemetry() here will panic with a nil pointer dereference!
```

---

## 🛠️ Today's Coding Exercise
Go to the [01_interfaces/main.go](file:///Users/clicknam/Documents/GitHub/go-scratch-pad/01_interfaces/main.go) file.
Your task is to:
1. Implement the `Megapack` device (pointer receiver).
2. Implement the `SolarInverter` device (value receiver).
3. Implement `GetTotalPower(devices []Device) float64` which aggregates their power and handles errors gracefully.
4. Run code and answer the conceptual questions in the comments.
