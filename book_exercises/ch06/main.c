#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

typedef struct s_person {
    char* name;
    char* surname;
    int age;
} t_person;

#define ITER 100000000

t_person* allocate_person(t_person* person);
void free_person_list(t_person** list);

int main(void)
{
    t_person** list = (t_person**)malloc(sizeof(t_person*) * ITER);
    if (!list) {
        perror("malloc");
        exit(EXIT_FAILURE);
    }
    t_person* person_list = (t_person*)malloc(sizeof(t_person) * ITER);
    t_person* head_person_list = person_list;
    for (int i = 0; i < ITER; i++) {
        list[i] = allocate_person(person_list);
        person_list++;
    }
    printf("%s\n", list[0]->name);
    free(list);
    free(head_person_list);
    // free_person_list(list);
    return (0);
}

t_person* allocate_person(t_person* person)
{
    if (!person) {
        perror("malloc");
        exit(EXIT_FAILURE);
    }
    person->name = "bob";
    person->surname = "smith";
    person->age = 42;
    return person;
}

void free_person_list(t_person** list)
{
    for (int i = 0; i < ITER; i++) {
        free(list[i]);
    }
    free(list);
}
