import random
import yaml


def get_workout(number_of_musclegroups=7):
    with open('./exercises.yaml') as f:
        thugletics = yaml.safe_load(f)
    musclegroups = random.sample(thugletics, number_of_musclegroups)
    return (random.choice(thugletics[musclegroup]) for musclegroup in musclegroups)


if __name__ == "__main__":
    for exercise in get_workout():
        print(exercise)
