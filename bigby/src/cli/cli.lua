-- local character = require("character")
-- local monster = require("monster")
-- local npc = require("npc")

function MainTitle()
	print("")
	print("--------------------------------------------------")
	print("OWLBEAR")
	print("A cli app for handling dungeons and dragons stuff")
	print("--------------------------------------------------")
	print("")
end

function MainHelp()
	MainTitle()

	print("USAGE: owlbear [COMMAND] [OPTIONS]")
	print("COMMAND:")
	print("  character   character command")
	print("  npc         npc command")
	print("  monster     monster command")
	print("OPTIONS:")
	print("  --help      show this message")
end

function Main(args)
	if #args < 1 then
		MainHelp()
		return
	end

	if args[1] == "character" then
		CharacterHandler(args)
		--   and args[2] == "character" then
		-- local name = nil
		--
		-- for i = 3, #args do
		-- 	if args[i] == "--name" then
		-- 		if i + 1 <= #args then
		-- 			name = args[i + 1]
		-- 		else
		-- 			print("Erro: --name requer um valor.")
		-- 			return
		-- 		end
		-- 	elseif args[i] == "--help" then
		-- 		MainHelp()
		-- 		return
		-- 	else
		-- 		print("Erro: opção desconhecida " .. args[i])
		-- 		return
		-- 	end
		-- end
		--
		-- if name then
		-- 	print("Character and stuff")
		-- 	-- create_character(name)
		-- else
		-- 	print("Erro: Nome do personagem não fornecido.")
		-- end
	elseif args[1] == "npc" then
		NpcHandler(args)
	elseif args[1] == "monster" then
		MonsterHandler(args)
	else
		MainHelp()
	end
end

-- Call the main function
Main(arg)
