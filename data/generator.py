import random
import csv


def generate_random_data(num_records, filename="bulk_data.csv"):
    names = ["Alice", "Bob", "Charlie", "David", "Eva", "Frank", "Grace", "Helen", "Ivy", "Jack",
             "Kate", "Liam", "Mona", "Nina", "Oscar", "Paul", "Quinn", "Rose", "Steve", "Tina", "Alireza",
             "Ali", "Reza", "Hossein", "Mohammad", "Sara", "Fatemeh", "Zahra", "Maryam", "Narges", "Mahdi",
             "Amir", "Sina", "Parsa", "Shayan", "Arman", "Yasaman", "Elham", "Fereshteh", "Hassan", "Ramin",
             "Shirin", "Parvin", "Nasrin", "Farhad", "Behnam", "Sepideh", "Bahram", "Pouya", "Golnaz", "Roya",
             "Niloufar", "Atiyeh", "Sahar", "Masoud", "Azadeh", "Farzaneh", "Kaveh", "Neda", "Sohrab", "Milad",
             "Mehdi", "Kamran", "Shiva", "Ladan", "Negar", "Setareh", "Katayoun", "Samira", "Fereshte", "Afshin",
             "Mehrdad", "Dariush", "Arash", "Yalda", "Shahram", "Ehsan", "Shadi", "Leila", "Mojtaba", "Homa",
             "Azar", "Shabnam", "Babak", "Siavash", "Fariba", "Mina", "Pegah", "Javad", "Vahid", "Elnaz",
             "Mandana", "Rasoul", "Navid", "Hamid", "Ghazal", "Shervin", "Dorsa", "Nasim", "Elaheh", "Soraya",
             "Ziba", "Anahita", "Kiana", "Shahab", "Simin", "Pouneh", "Younes", "Mohsen", "Taraneh", "Goli",
             "Nader", "Keyvan", "Farshid", "Shahrzad", "Negin", "Vanda", "Morteza", "Behzad", "Bahareh"]

    with open(filename, mode='w', newline='') as file:
        writer = csv.writer(file)
        writer.writerow(["id", "name", "grade"])

        # Write data rows
        for i in range(1, num_records + 1):
            name = random.choice(names)

            grade = round(random.uniform(0, 20), 2)
            writer.writerow([i, name, grade])

    print(f"Generated {num_records} records in {filename}")


generate_random_data(100000,"chemistry.csv")
