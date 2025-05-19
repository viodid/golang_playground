#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

typedef struct s_person {
    char* name;
    char* surname;
    int age;
} t_person;

#define ITER 100000000

t_person* allocate_person(void);
void free_person_list(t_person** list);

int main(void)
{
    t_person** list = (t_person**)malloc(sizeof(t_person*) * ITER);
    if (!list) {
        perror("malloc");
        exit(EXIT_FAILURE);
    }
    for (int i = 0; i < ITER; i++) {
        list[i] = allocate_person();
    }
    printf("%s\n", list[0]->name);
    free_person_list(list);
    return (0);
}

t_person* allocate_person(void)
{
    t_person* person = (t_person*)malloc(sizeof(t_person));
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
        t_person* p = list[i];
        free(p);
    }
    free(list);
}
