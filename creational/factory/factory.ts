// Factory Method is a creational design pattern that provides an interface for creating objects
// but allows subclasses to alter the type of objects that will be created.

// Shape interface defines the operations that concrete shapes must implement.
interface Shape {
  draw(): string;
}

// Concrete Circle class implements the Shape interface.
class Circle implements Shape {
  draw(): string {
    return "Drawing a circle";
  }
}

// Concrete Square class implements the Shape interface.
class Square implements Shape {
  draw(): string {
    return "Drawing a square";
  }
}

// Abstract ShapeCreator class defines the factory method for creating shapes.
abstract class ShapeCreator {
  abstract createShape(): Shape;

  create(): string {
    const shape = this.createShape();
    return `ShapeCreator: Created shape - ${shape.draw()}`;
  }
}

// Concrete CircleCreator class extends ShapeCreator to create circles.
class CircleCreator extends ShapeCreator {
  // here create shape is factory method
  createShape(): Shape {
    return new Circle();
  }
}

// Concrete SquareCreator class extends ShapeCreator to create squares.
class SquareCreator extends ShapeCreator {
  // here create shape is factory method
  createShape(): Shape {
    return new Square();
  }
}

// Client code can use the creators to create shapes without knowing the concrete classes.
function clientCode(creator: ShapeCreator) {
  console.log(creator.create());
}

// Application code:

console.log("App: Drawing a circle.");
clientCode(new CircleCreator());
console.log("");

console.log("App: Drawing a square.");
clientCode(new SquareCreator());
