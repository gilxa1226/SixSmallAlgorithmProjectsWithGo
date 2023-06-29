package main

import "fmt"

const num_disks = 3

func push(post []int, disk int) []int {
	return append([]int{disk}, post...)
}

func pop(post []int) (int, []int) {
	return post[0], post[1:]
}

func move_disk(posts [][]int, from_post int, to_post int) {
	var data int

	data, posts[from_post] = pop(posts[from_post])
	posts[to_post] = push(posts[to_post], data)
}

func draw_posts(posts [][]int) {
	for p := 0; p < len(posts); p++ {
		if len(posts[p]) < num_disks {
			for d := len(posts[p]); d < num_disks; d++ {
				posts[p] = push(posts[p], 0)
			}
		}
	}

	for d := 0; d < num_disks; d++ {
		for p := 0; p < num_disks; p++ {
			fmt.Printf("%2d", posts[p][d])
		}
		fmt.Println()
	}

	for p := 0; p < num_disks; p++ {
		for d := 0; d < num_disks; d++ {
			if len(posts[p]) > 0 && posts[p][0] == 0 {
				_, posts[p] = pop(posts[p])
			}
		}
	}
}

func move_disks(posts [][]int, num_to_move, from_post, to_post, temp_post int) {
	if num_to_move > 1 {
		move_disks(posts, num_to_move-1, from_post, temp_post, to_post)
	}

	move_disk(posts, from_post, to_post)
	draw_posts(posts)
	fmt.Println(" -----")

	if num_to_move > 1 {
		move_disks(posts, num_to_move-1, temp_post, to_post, from_post)
	}
}

func main() {
	// Make three posts.
	posts := [][]int{}

	// Push the disks onto post 0 biggest first.
	posts = append(posts, []int{})
	for disk := num_disks; disk > 0; disk-- {
		posts[0] = push(posts[0], disk)
	}

	// Make the other posts empty.
	for p := 1; p < 3; p++ {
		posts = append(posts, []int{})
	}

	// Draw the initial setup.
	draw_posts(posts)
	fmt.Println(" -----")

	// Move the disks.
	move_disks(posts, num_disks, 0, 1, 2)
}
