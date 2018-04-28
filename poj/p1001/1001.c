#include<stdio.h>
#include<stdlib.h>
#include<string.h>

const int MAX_BASE_LEN = 7;
const int MAX_RET_LEN = 155;
const char *point = ".";
char err_msg[100];

int standardize(const char *src, char *dest, int *bits);
int str_pow(char *str, int exp, char *result);
int debug_output(char *str);
int multi_by_single_num(char *x, char y, char *ret);
int combine_2_str_num(char *target, char *delta, int exp);

/**
 * standardize input str_num, calculate effective number of fractional digits
 * preparing for str_pow
 *
 * @param char *src    source string
 * @param char *dest   destination string
 * @param int  *bits   effective number of fractional digits  (after trimming trailing '0')
 *
 * @sample:
 * 95.123 ====> 95123, 3
 * 0.4321 ====> 4321,  4
 * 1.0100 ====> 101,   2
 * .00001 ====> 1,     5
 * 000010 ====> 10,    0
 * 000000 ====> 0,     0
 * 000.00 ====> 0,     0
 * .00000 ====> 0,     0
 * 1.0000 ====> 1,     0
 */
int standardize(const char *src, char *dest, int *bits) {
    char *dest_bg = dest; /* save point of dest string */
    size_t len = strlen(src);
    *bits = 0;
    char *src_copy = malloc(sizeof(src));
    memcpy(src_copy, src, len);

    /* trim trailing '0' */
    char *ptr = strstr(src, point);
    if (ptr) {
        int i;
        for(i = len - 1; i >= 0; i--) {
            if (src_copy[i] != '0') {
                break;
            }
        }
        if (i < len - 1) {
            src_copy[i+1] = '\0';
        }
        /* calculate effective number of franctional digits */
        for (; src_copy[i] != '.'; i--) {
            (*bits)++;
        }
        len = strlen(src_copy);
    }

    /* copy char except '.' */
    int zero_flag = 0;
    int non_zero = 0;
    for(int i = 0; i < len; i++) {
        if (src_copy[i] != '0' && zero_flag == 0) {
            zero_flag = 1;
        }
        if (zero_flag == 1 && src_copy[i] != '.') {
            non_zero = 1;
            *dest++ = src_copy[i];
        }
    }
    free(src_copy);

    /* trim leading '0' */
    if (*dest_bg == '0') {
        char *cur = dest_bg;
        while(*cur == '0' && cur != dest) {
            cur++;
        }
        if (cur != dest) {
            while(cur != dest) {
                *dest_bg++ = *cur++;
            }
            *dest_bg = '\0';
        } else {
            non_zero = 0;
        }
    }

    /* set dest to '0' if all digits has been filtered */
    if (non_zero == 0) {
        *dest++ = '0';
    }
    *dest = '\0';
    return 0;
}

/**
 * debug func, output string after trimming leading '0'
 */
int debug_output(char* str) {
    int len = strlen(str);
    int zero_flag = 0;
    for(int i = 0; i < len; i++) {
        if(str[i] != '0' && zero_flag == 0) {
            zero_flag = 1;
        }
        if(zero_flag == 1) {
            printf("%c", str[i]);
        }
    }
    printf("\n");
    return 0;
}

/**
 * multiply string by char
 *
 * @param char *x    source string, multiplicand
 * @param char y     source char, multiplier
 * @param char *ret  destination string
 */
int multi_by_single_num(char *x, char y, char *ret) {
    /*printf("[multi_by_single_num] x:%s, y:%c\n", x, y);*/
    /*printf("[multi_by_single_num] ret:%p\n", ret);*/

    size_t len = strlen(x);

    int carry = 0;
    int y_t = y - '0';
    for(int i = len - 1; i >= 0; i--) {
        /*printf("[multi_by_single_num] loop i:%d, x[i]:%c\n", i, x[i]);*/
        int x_i = x[i] - '0';
        int tmp = x_i * y_t + carry;
        ret[i] = tmp%10 + '0';
        carry = tmp / 10;
        /*printf("[multi_by_single_num] tmp:%d, carry:%d, ret:%s\n", tmp, carry, ret);*/
    }
    if (carry > 0) {
        printf("[multi_by_single_num] multiplication overflow, x:%s, y:%c", x, y);
        /*ret[0] = carry + '0';*/
        strcpy(err_msg, "[multi_by_single_num] multiplication overflow.");
    }
    /*printf("[multi_by_single_num] finally, ret:%s\n", ret);*/
    return 0;
}

/**
 * combine multiplication result string
 *
 * @param char *target    destination string, contains result of last step
 * @param char *delta     source char, result of this step
 * @param int  exp        digits to left shit when combining two string
 */
int combine_2_str_num(char *target, char *delta, int exp) {
    /*printf("[combine_2_str_num] target:%s, delta:%s, exp:%d\n", target, delta, exp);*/
    size_t len_target = strlen(target);
    size_t len_delta = strlen(delta);

    int carry = 0;
    int m_num = len_target - len_delta - exp;
    /*printf("[combine_2_str_num] len_target:%zu, len_delta:%zu, m_num:%d\n", len_target, len_delta, m_num);*/
    for (int i = len_target - 1; i >= 0; i--) {
        if (len_target - exp >= i && i > m_num) {
            int tmp = (target[i] - '0') + (delta[ i - m_num - 1] - '0') + carry;
            target[i] = tmp % 10 + '0';
            carry = tmp / 10;
            /*printf("[combine_2_str_num] loop if 1, i:%d, tmp:%d, target[i]:%c, carry:%d\n", i, tmp, target[i], carry);*/
        }
        if (i == m_num) {
            int tmp = target[i] - '0' + carry;
            target[i] = tmp % 10 + '0';
            carry = tmp / 10;
            /*printf("[combine_2_str_num] loop if 2, i:%d, tmp:%d, target[i]:%c, carry:%d\n", i, tmp, target[i], carry);*/
        }
        if (i == m_num - 1 && carry > 0) {
            target[i] = carry + '0';
            carry = 0;
        }
    }

    /* debug */
    /*printf("combine_2_str_num debug:");*/
    /*debug_output(target);*/
    return 0;
}

int multi_by_str(char *x, char *y, char *ret) {
    /*printf("[multi_by_str] x:%s, y:%s, ret:%s\n", x, y, ret);*/

    size_t len = strlen(y);
    char tmp[MAX_RET_LEN];
    for(int i = len - 1; i >= 0; i--) {
        memset(tmp, '0', MAX_RET_LEN);
        tmp[MAX_RET_LEN - 1 ] = '\0';
        multi_by_single_num(x, y[i], tmp);

        /*printf("i:%d, y[i]:%c, x:", i, y[i]);*/
        /*debug_output(x);*/
        /*printf("tmp:");*/
        /*debug_output(tmp);*/

        combine_2_str_num(ret, tmp, len - i);
        /*printf("ret:");*/
        /*debug_output(ret);*/
    }
    /*debug_output(ret);*/
    return 0;
}

int format_result(char *str_num, int bits, char *ret) {
    /*printf("[format_and_output] str_num:%s\n", str_num);*/
    int len = strlen(str_num);
    if (bits > 0) {
        if (str_num[0] != '0') {
            strcpy(err_msg, "[format_and_output] left shif overflow");
            return 1;
        }

        for(int i = 1; i < len - bits; i++) {
            str_num[ i - 1 ] = str_num[i];
        }
        str_num[len - bits - 1] = '.';
    }

    int zero_flag = 0;
    int counter = 0;
    for(int i = 0; i < len; i++) {
        /* skip leading '0' */
        if(str_num[i] != '\0' && str_num[i] != '0' && zero_flag == 0)  {
            zero_flag = 1;
        }
        if (zero_flag == 1) {
            ret[counter++] = str_num[i];
            /*printf("%c", str_num[i]);*/
        }
    }
    if(zero_flag == 0) {
        ret[counter++] = '0';
    }
    
    /* trimming trailing '0' */
    if (bits > 0) {
        /* get index of last effective digit */
        int i;
        for (i = counter - 1; i >= 0; i--) {
            if(ret[i] != '0' && ret[i] != '.') {
                break;
            }
        }
        if (i == 0) {
            ret[0] = '0';
        }

        ret[i + 1] = '\0';
    } else {
        ret[counter] = '\0';
    }
    return 0;
}

int str_pow(char *str, int exp, char *result) {
    char *stdd = malloc(MAX_BASE_LEN);
    int bits = 0;
    standardize(str, stdd, &bits);

    if (stdd[0] == '0' && strlen(stdd) == 1 && bits == 0) {
        result[0] = '0';
        result[1] = '\0';
        return 0;
    }

    if (exp == 0) {
        result[0] = '1';
        result[1] = '\0';
        return 0;
    }

    char *cur = malloc(MAX_RET_LEN);
    memset(cur, '0', MAX_RET_LEN);
    cur[MAX_RET_LEN - 1] = '\0';

    strcpy(&cur[MAX_RET_LEN - strlen(stdd) - 1], stdd);
    int total_bits = bits * exp;
    /*printf("base:%s, stdd:%s, bits:%d, total_bits:%d, ", str, stdd, bits, total_bits);*/

    for (int i = exp - 1; i > 0 ; i--) {
        memset(result, '0', MAX_RET_LEN);
        result[MAX_RET_LEN - 1] = '\0';

        multi_by_str(cur, stdd, result);
        strcpy(cur, result);
    }

    format_result(cur, total_bits, result);
    return 0;
}

int main() {
    char base[MAX_BASE_LEN];
    int exp;

    char *ret = malloc(MAX_RET_LEN);
    while(scanf("%s %d", base, &exp) != EOF) {
        memset(ret, '0', MAX_RET_LEN);
        str_pow(base, exp, ret);
        /*printf("%s ^ %d ====> %s\n", base, exp, ret);*/
        printf("%s\n", ret);
        /*debug_output(ret);*/

        /* debug */
        /*char *stdd = malloc(MAX_BASE_LEN);*/
        /*int b = 0;*/
        /*standardize(base, stdd, &b);*/
        /*printf("%s ==== %s, bits:%d\n", base, stdd, b);*/
    }
    return 0;
}
