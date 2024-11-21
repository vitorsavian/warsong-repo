function NpcHelp()
	print("Usage: owlbear npc [COMMAND] [OPTIONS]")
	print("COMMAND:")
	print("  create     Create npc")
	print("  update     Update npc")
	print("  get        Get npc")
	print("  delete     Delete npc")
	print("OPTIONS:")
	print("  --help     Show this message")
end

function NpcCreateHelp()
	print("Usage: owlbear npc create [OPTIONS]")
	print("COMMAND:")
	print("")
	print("OPTIONS:")
	print("  --help     Show this message")
end

function NpcUpdateHelp()
	print("Usage: owlbear npc update [OPTIONS]")
	print("COMMAND:")
	print("")
	print("OPTIONS:")
	print("  --help     Show this message")
end

function NpcGetHelp()
	print("Usage: owlbear npc get [OPTIONS]")
	print("COMMAND:")
	print("")
	print("OPTIONS:")
	print("  --help     Show this message")
end

function NpcDeleteHelp()
	print("Usage: owlbear npc delete [OPTIONS]")
	print("COMMAND:")
	print("")
	print("OPTIONS:")
	print("  --help     Show this message")
end

function NpcHandler(args)
	for i = 2, #args do
		if args[i] == "--help" then
			NpcHelp()
			return
		end
	end
end
