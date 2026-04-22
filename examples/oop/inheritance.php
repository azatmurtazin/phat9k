<?php

class Animal {
    protected $name;

    public function __construct($name) {
        $this->name = $name;
    }

    public function speak() {
        return "Some sound";
    }
}

class Dog extends Animal {
    public function speak() {
        return $this->name . " says Woof!";
    }
}

class Cat extends Animal {
    public function speak() {
        return $this->name . " says Meow!";
    }
}

$dog = new Dog("Buddy");
$cat = new Cat("Whiskers");

echo $dog->speak() . "\n";
echo $cat->speak() . "\n";