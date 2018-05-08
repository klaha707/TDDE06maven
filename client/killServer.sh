#!/bin/bash
kill -9 $(ps ux | grep [s]erver- | awk '{print $2}')
