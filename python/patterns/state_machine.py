class StateMachine:
    def __init__(self) -> None:
        self.state = "INIT"

    def process_input(self, user_input) -> None:
        if self.state == "INIT":
            print("Welcome! Please enter your name:")
            self.state = "WAITING_FOR_NAME"
        elif self.state == "WAITING_FOR_NAME":
            print(f"Hello, {user_input}!")
            print("Do you want to proceed? (yes/no)")
            self.state = "WAITING_FOR_CONFIRMATION"
        elif self.state == "WAITING_FOR_CONFIRMATION":
            if user_input.lower() == "yes":
                print("Let's proceed to the next step.")
                self.state = "NEXT_STEP"
            elif user_input.lower() == "no":
                print("Goodbye!")
                self.state = "END"
            else:
                print("Please answer 'yes' or 'no'.")

    def run(self) -> None:
        while self.state != "END":
            user_input: str = input("> ")
            self.process_input(user_input)


if __name__ == "__main__":
    state_machine = StateMachine()
    state_machine.run()
