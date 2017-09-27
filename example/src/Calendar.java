/******************************************************************************
 *  Compilation:  javac Calendar.java
 *  Execution:    java Calendar month year
 *
 *  This program takes the month and year as command-line arguments
 *  and prints a calendar for that month.
 * 
 *  % java Calendar 7 2005
 *   July 2005
 *   S  M  T  W Th  F  S
 *                  1  2 
 *   3  4  5  6  7  8  9  
 *  10 11 12 13 14 15 16 
 *  17 18 19 20 21 22 23 
 *  24 25 26 27 28 29 30 
 *  31 
 *  
 ******************************************************************************/

public class Calendar {

   /***************************************************************************
    *  Given the month, day, and year, return which day
    *  of the week it falls on according to the Gregorian calendar.
    *  For month, use 1 for January, 2 for February, and so forth.
    *  Returns 0 for Sunday, 1 for Monday, and so forth.
    ***************************************************************************/
    public static int day(int month, int day, int year) {
        int y = year - (14 - month) / 12;
        int x = y + y/4 - y/100 + y/400;
        int m = month + 12 * ((14 - month) / 12) - 2;
        int d = (day + x + (31*m)/12) % 7;
        return d;
    }

    // return true if the given year is a leap year
    public static boolean isLeapYear(int year) {
        if  ((year % 4 == 0) && (year % 100 != 0)) return true;
        if  (year % 400 == 0) return true;
        return false;
    }

    public static void main(String[] args) {

        if(args.length < 2) throw new IllegalArgumentException("Date format -> mm yyyy");

        int month = Integer.parseInt(args[0]);    // month (Jan = 1, Dec = 12)
        int year = Integer.parseInt(args[1]);     // year

        // months[i] = name of month i
        String[] months = {
            "",                               // leave empty so that months[1] = "January"
            "January", "February", "March",
            "April", "May", "June",
            "July", "August", "September",
            "October", "November", "December"
        };

        // days[i] = number of days in month i
        int[] days = {
            0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31
        };

        // check for leap year
        if (month == 2 && isLeapYear(year)) days[month] = 29;


        // print calendar header
        System.out.println("   " + months[month] + " " + year);
        System.out.println(" S  M Tu  W Th  F  S");

        // starting day
        int d = day(month, 1, year);

        // print the calendar
        for (int i = 0; i < d; i++)
            System.out.print("   ");
        for (int i = 1; i <= days[month]; i++) {
            System.out.printf("%2d ", i);
            if (((i + d) % 7 == 0) || (i == days[month])) System.out.println();
        }

    }
}