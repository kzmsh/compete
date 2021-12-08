#include <bits/stdc++.h>

using namespace std;

using i32 = int32_t;
using i64 = int64_t;
using u32 = uint32_t;
using u64 = uint64_t;
using f32 = float;
using f64 = double;

template<typename T>
bool choose_min(T &min, const T &value) {
    if (min > value) {
        min = value;
        return true;
    }
    return false;
}

template<typename T>
bool choose_max(T &max, const T &value) {
    if (max < value) {
        max = value;
        return true;
    }
    return false;
}

template<typename T, typename = enable_if_t<is_integral_v<T>>>
bool is_prime(const T &integer) {
    if (integer == 2) {
        return true;
    }
    if (integer <= 1 || integer % 2 == 0) {
        return false;
    }
    for (int v = 3; v <= sqrt(integer); v += 2) {
        if (integer % v == 0) {
            return false;
        }
    }
    return true;
}

void run();

int main() {
    cin.tie(0);
    ios::sync_with_stdio(false);
    run();
}

i64 solve(vector<i64> A) {
    vector<i64> nums(0);
    unordered_map<i64, bool> appeared(0);
    for (const auto &a : A) {
        if (appeared[a]) {
            nums.push_back(a);
            appeared[a] = false;
        } else {
            appeared[a] = true;
        }
    }
    sort(nums.begin(), nums.end(), greater<i64>());
    return nums.size() < 2 ? 0 : nums[0] * nums[1];
}

void run() {
    i64 N;
    cin >> N;

    vector<i64> A(N);
    for (auto &a : A) {
        cin >> a;
    }

    cout << solve(A) << endl;
}
