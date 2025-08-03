import csv

input_file = 'student_data.csv'
output_file = 'cleaned_file.csv'

with open(input_file, 'r', newline='', encoding='utf-8') as infile:
    reader = list(csv.reader(infile))
    header = reader[0]
    expected_len = len(header)

    fixed_rows = [header]
    for index, row in enumerate(reader[1:], start=2):
        if len(row) == expected_len:
            fixed_rows.append(row)
        else:
            print(f"[❌] Row {index} has {len(row)} columns instead of {expected_len}: {row}")
            # Fix strategy: pad missing with empty strings, or truncate extra
            if len(row) < expected_len:
                row += [''] * (expected_len - len(row))
            else:
                row = row[:expected_len]
            print(f"[🔧] Fixed Row {index}: {row}")
            fixed_rows.append(row)

# Write cleaned CSV
with open(output_file, 'w', newline='', encoding='utf-8') as outfile:
    writer = csv.writer(outfile)
    writer.writerows(fixed_rows)

print(f"\n✅ Cleaned CSV saved to {output_file}")





# import csv

# input_file = 'DevOps_TA_copy.csv'
# cleaned_file = 'cleaned_file.csv'
# debug_file = 'debug_rows.csv'

# bad_rows = []

# with open(input_file, 'r', newline='', encoding='utf-8') as infile:
#     reader = list(csv.reader(infile))
#     header = reader[0]
#     expected_len = len(header)

#     fixed_rows = [header]

#     for index, row in enumerate(reader[1:], start=2):
#         row_len = len(row)

#         if row_len == expected_len:
#             fixed_rows.append(row)
#         else:
#             print(f"\n[❌] Row {index} is inconsistent:")
#             print(f"     ➤ Found {row_len} columns, expected {expected_len}")
#             print(f"     ➤ Raw Row: {row}")

#             bad_rows.append(
#                 [index, row_len, expected_len, "MISSING" if row_len < expected_len else "EXTRA", *row]
#             )

#             # Fix the row
#             if row_len < expected_len:
#                 row += [''] * (expected_len - row_len)
#             else:
#                 row = row[:expected_len]

#             print(f"[🔧] Fixed Row {index}: {row}")
#             fixed_rows.append(row)

# # Save cleaned file
# with open(cleaned_file, 'w', newline='', encoding='utf-8') as outfile:
#     writer = csv.writer(outfile)
#     writer.writerows(fixed_rows)

# # Save debug report
# if bad_rows:
#     with open(debug_file, 'w', newline='', encoding='utf-8') as debug_out:
#         debug_writer = csv.writer(debug_out)
#         debug_writer.writerow(["Row Number", "Actual Columns", "Expected Columns", "Issue Type", *header])
#         debug_writer.writerows(bad_rows)

#     print(f"\n⚠️  {len(bad_rows)} problematic rows found.")
#     print(f"🛠️  Debug info saved to '{debug_file}'")
# else:
#     print("\n✅ No issues found. All rows were consistent.")

# print(f"📄 Cleaned CSV saved to '{cleaned_file}'")
