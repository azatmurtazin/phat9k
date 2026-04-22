interface Drawable {
    public function draw(): string;
}

class Circle implements Drawable {
    public function draw(): string {
        return "Drawing a circle";
    }
}

class Square implements Drawable {
    public function draw(): string {
        return "Drawing a square";
    }
}

$shapes = [new Circle(), new Square()];

foreach ($shapes as $shape) {
    echo $shape->draw() . "\n";
}